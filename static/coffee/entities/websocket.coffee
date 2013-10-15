define ["devfeed", "golem"], (Devfeed, Golem) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    class Entities.WebSocket
      constructor: ->
        @WsConn = new Golem.Connection("ws://" + CONFIG.baseUrl + "/ws", CONFIG.wsDebug)
        @WsConn.on "open", @open
        @WsConn.on "message", @message
   
      open: =>
        @WsConn.emit "subscribe", channel: CONFIG.projChanName
   
      message: (data) =>
        console.log data

    Devfeed.addInitializer ->
      window.ws = new Entities.WebSocket()

  return Devfeed.Entities.WebSocket
