define ["devfeed"], (Devfeed) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    Entities.Chat = {}

    class Entities.Chat.Message extends Backbone.Model
      defaults:
        id: null
        author_id: null
        project_id: null
        body: null
        created: null

    class Entities.Chat.Messages extends Backbone.Model
      model: Entities.Chat.Message
      url: "/api/messages"
      comparator: (message) ->
        return message.get("id")

    messages = new Entities.Chat.Messages

    API =
      sendMessage: (projectId, body) ->
        messages.save project_id: projectId, body: body

    Devfeed.commands.setHandler "chat:message:send", (projectId, body) ->
      API.sendMessage(projectId, body)

  return Devfeed.Entities.Chat
