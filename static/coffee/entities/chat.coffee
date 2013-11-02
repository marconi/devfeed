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
      fetchMessages: (projId) ->
        defer = $.Deferred()
        messages.fetch
          reset: true
          data:
            project_id: projId
          success: (collection, response, options) ->
            defer.resolve messages
          error: (collection, response, options) ->
            defer.resolve null
        return defer.promise()

      sendMessage: (objId, body) ->
        messages.create project_id: objId, body: body

    Devfeed.reqres.setHandler "chat:messages:fetch", (objId) ->
      return API.fetchMessages(objId)

    Devfeed.commands.setHandler "chat:message:send", (projId, body) ->
      API.sendMessage(projId, body)

  return Devfeed.Entities.Chat
