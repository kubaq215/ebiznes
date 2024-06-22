package com.example

import dev.kord.common.entity.Snowflake
import dev.kord.core.Kord
import dev.kord.core.Kord.*
import dev.kord.rest.builder.interaction.*
import dev.kord.core.on
import dev.kord.core.event.interaction.*
import dev.kord.core.behavior.interaction.*
import dev.kord.rest.builder.message.modify.*
import dev.kord.core.entity.interaction.response.*
import dev.kord.core.behavior.interaction.response.*
import io.ktor.http.*
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.plugins.HttpTimeout
import io.ktor.client.statement.*
import io.ktor.client.request.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import io.ktor.server.request.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.*
import kotlinx.serialization.json.Json
import kotlinx.coroutines.runBlocking
import io.github.cdimascio.dotenv.dotenv

@Serializable
data class ChatRequest(val user_id: String, val message: String)

@Serializable
data class ChatResponse(val response: String)

val client = HttpClient(CIO) {
    install(ContentNegotiation) {
        json(Json {
            prettyPrint = true
            isLenient = true
        })
    }
    install(HttpTimeout) {
        requestTimeoutMillis = 90_000
    }
}

suspend fun handleLLMQuery(userId: String, query: String): ChatResponse {
    val chatRequest = ChatRequest(user_id = userId, message = query)
    return try {
        val response: HttpResponse = client.post("http://localhost:5000/chat") {
            contentType(ContentType.Application.Json)
            setBody(chatRequest)
        }
        Json.decodeFromString<ChatResponse>(response.bodyAsText())
    } catch (e: Exception) {
        ChatResponse("There was an error processing your request. Please try again later.")
    }
}

// fun sendMessage(message: String, webhookUrl: String) = runBlocking {
//     val response = client.post(webhookUrl) {
//         contentType(ContentType.Application.Json)
//         setBody(mapOf("content" to message))
//     }
//     println("Response status: " + response.status)
// }

suspend fun main() {
    val dotenv = dotenv()
    val discordToken = dotenv["DISCORD_BOT_TOKEN"]
    val webhookUrl = dotenv["DISCORD_WEBHOOK_URL"]

    if (discordToken == null || webhookUrl == null) {
        println("DISCORD_BOT_TOKEN and DISCORD_WEBHOOK_URL environment variables must be set")
        return
    }

    val kord = Kord(discordToken)

    kord.createGuildChatInputCommand(
        Snowflake(1253272703222419466),
        "ask",
        "Ask about the store or items"
    ) {
        string("text", "Your question") {
            required = true
        }
    }

    kord.on<GuildChatInputCommandInteractionCreateEvent> {
        val response = interaction.deferPublicResponse()
        val command = interaction.command
        val query = command.strings["text"]!! // it's required so it's never null
        val userId = interaction.user.id.toString()
        println("User ID: $userId")
        val llmAnswer = handleLLMQuery(userId, query).response
        val res_content = "Q: $query\nA: $llmAnswer"
        response.respond { content = res_content }
    }

    kord.login()
    
    // embeddedServer(Netty, port = 8080) {
    //     routing {
    //         post("/llm-query") {
    //             println("Received request")
    //             val request = call.receiveText()
    //             val chatRequest = Json.decodeFromString<ChatRequest>(request)
                
    //             val userId = chatRequest.user_id
    //             val query = chatRequest.message

    //             println("User ID: $userId")
    //             println("Query: $query")

    //             if (userId.isBlank()) {
    //                 call.respond(HttpStatusCode.BadRequest, "user_id is required")
    //                 return@post
    //             }

    //             val llmAnswer = handleLLMQuery(userId, query)
    //             println("Response: " + llmAnswer.response)
    //             sendMessage(llmAnswer.response, webhookUrl)
    //             call.respond(llmAnswer.response)
    //         }
    //         get("/health") {
    //             call.respond(HttpStatusCode.OK, "OK")
    //         }
    //     }
    // }.start(wait = true)
}
