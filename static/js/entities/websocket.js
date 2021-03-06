(function() {
  var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

  define(["devfeed", "golem"], function(Devfeed, Golem) {
    Devfeed.module("Entities", function(Entities, Devfeed, Backbone, Marionette, $, _) {
      var API, websocket;
      Entities.WebSocket = (function() {
        function WebSocket() {
          this.projectSubscribe = __bind(this.projectSubscribe, this);
          this.projectSynced = __bind(this.projectSynced, this);
          this.message = __bind(this.message, this);
          this.open = __bind(this.open, this);
          this.wsConn = new Golem.Connection("ws://" + CONFIG.baseUrl + "/ws", CONFIG.wsDebug);
          this.wsConn.on("open", this.open);
          this.wsConn.on("message", this.message);
          this.wsConn.on("project:synced", this.projectSynced);
        }

        WebSocket.prototype.open = function() {
          var userSession;
          userSession = Devfeed.request("user:session");
          return this.wsConn.emit("init", {
            user_id: userSession.get("id")
          });
        };

        WebSocket.prototype.message = function(message) {
          return console.log(message);
        };

        WebSocket.prototype.projectSynced = function(projectId) {
          return Devfeed.trigger("project:synced", projectId);
        };

        WebSocket.prototype.projectSubscribe = function(projectId) {
          return this.wsConn.emit("project:subscribe", {
            project_id: projectId
          });
        };

        return WebSocket;

      })();
      websocket = null;
      API = {
        createWebSocket: function() {
          if (!websocket) {
            return websocket = new Entities.WebSocket();
          }
        },
        getWebSocket: function() {
          return websocket;
        },
        projectSubscribe: function(projectId) {
          return websocket.projectSubscribe(projectId);
        }
      };
      Devfeed.on("loggedin", function() {
        return API.createWebSocket();
      });
      Devfeed.reqres.setHandler("websocket:entity", function() {
        return API.getWebSocket();
      });
      return Devfeed.commands.setHandler("ws:project:subscribe", function(projectId) {
        return API.projectSubscribe(projectId);
      });
    });
    return Devfeed.Entities.WebSocket;
  });

}).call(this);
