define ["devfeed", "common_model"], (Devfeed, CommonModel) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    Entities.Chat = {}

    class Entities.Chat.Message extends CommonModel.BaseModel
      defaults:
        id: null
        author:
          id: null
          name: null
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
        defer = $.Deferred()
        # TODO: checkout if collection.create has a flag to not add the
        # model into collection until the request is done since it won't
        # render all the fields. 
        message = new Entities.Chat.Message
          project_id: objId, body: body
        message.save null,
          success: ->
            messages.add(message)
            defer.resolve message
          error: ->
            defer.resolve null
        return defer.promise()

    Devfeed.reqres.setHandler "chat:messages:fetch", (objId) ->
      return API.fetchMessages(objId)

    Devfeed.reqres.setHandler "chat:message:send", (projId, body) ->
      return API.sendMessage(projId, body)

  return Devfeed.Entities.Chat
