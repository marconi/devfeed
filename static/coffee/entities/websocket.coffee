define ["devfeed", "golem"], (Devfeed, Golem) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    class Entities.WebSocket
      constructor: ->
        @WsConn = new Golem.Connection("ws://" + CONFIG.baseUrl + "/ws", CONFIG.wsDebug)
        @WsConn.on "open", @open
        @WsConn.on "project:synced", @projectSynced
   
      open: =>
        userSession = Devfeed.request("user:session")
        @WsConn.emit "init", user_id: userSession.get("id")
   
      projectSynced: (projectId) =>
        Devfeed.trigger "project:synced", projectId

    websocket = null

    API =
      createWebSocket: ->
        # only create when we don't have one yet so we can call this
        # api multiple times without creating each time.
        if not websocket
          websocket = new Entities.WebSocket()
      getWebSocket: ->
        return websocket

    Devfeed.on "loggedin", ->
      # create websocket when someone logs-in
      API.createWebSocket()

    Devfeed.reqres.setHandler "websocket:entity", ->
      return API.getWebSocket()

    Devfeed.commands.setHandler "websocket:create", ->
      API.createWebSocket()

  return Devfeed.Entities.WebSocket
