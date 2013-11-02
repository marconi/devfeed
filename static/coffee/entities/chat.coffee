define ["devfeed", "common_model"], (Devfeed, CommonModel) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    Entities.Chat = {}

    class Entities.Chat.Message extends CommonModel.BaseModel
      defaults:
        id: null
        author_id: null
        project_id: null
        body: null
        created: null
      url: "/api/messages"

    class Entities.Chat.Messages extends CommonModel.BaseCollection
      model: Entities.Chat.Message
      url: "/api/messages"

    messages = new Entities.Chat.Messages

    API =
      sendMessage: (objId, body) ->
        messages.create project_id: objId, body: body
        console.log messages

    Devfeed.commands.setHandler "chat:message:send", (objId, body) ->
      API.sendMessage(objId, body)

  return Devfeed.Entities.Chat
