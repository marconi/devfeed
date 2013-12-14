define ["devfeed", "golem"], (Devfeed, Golem) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    class Entities.WebSocket
      constructor: ->
        @wsConn = new Golem.Connection("ws://" + CONFIG.baseUrl + "/ws", CONFIG.wsDebug)
        @wsConn.on "open", @open
        @wsConn.on "message", @message
        @wsConn.on "project:synced", @projectSynced
   
      open: =>
        userSession = Devfeed.request("user:session")
        @wsConn.emit "init", user_id: userSession.get("id")
   
      message: (message) =>
        console.log message

      projectSynced: (projectId) =>
        Devfeed.trigger "project:synced", projectId

      projectSubscribe: (projectId) =>
        @wsConn.emit "project:subscribe", project_id: projectId

    websocket = null

    API =
      createWebSocket: ->
        # only create when we don't have one yet so we can call this
        # api multiple times without creating each time.
        if not websocket
          websocket = new Entities.WebSocket()

      getWebSocket: ->
        return websocket

      projectSubscribe: (projectId) ->
        websocket.projectSubscribe(projectId)

    Devfeed.on "loggedin", ->
      # create websocket when someone logs-in
      API.createWebSocket()

    Devfeed.reqres.setHandler "websocket:entity", ->
      return API.getWebSocket()

    Devfeed.commands.setHandler "ws:project:subscribe", (projectId) ->
      API.projectSubscribe(projectId)

  return Devfeed.Entities.WebSocket
