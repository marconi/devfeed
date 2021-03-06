(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "common_model"], function(Devfeed, CommonModel) {
    Devfeed.module("Entities", function(Entities, Devfeed, Backbone, Marionette, $, _) {
      var API, messages, _ref, _ref1;
      Entities.Chat = {};
      Entities.Chat.Message = (function(_super) {
        __extends(Message, _super);

        function Message() {
          _ref = Message.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Message.prototype.defaults = {
          id: null,
          author: {
            id: null,
            name: null
          },
          project_id: null,
          body: null,
          created: null
        };

        Message.prototype.url = "/api/messages";

        return Message;

      })(CommonModel.BaseModel);
      Entities.Chat.Messages = (function(_super) {
        __extends(Messages, _super);

        function Messages() {
          _ref1 = Messages.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        Messages.prototype.model = Entities.Chat.Message;

        Messages.prototype.url = "/api/messages";

        return Messages;

      })(CommonModel.BaseCollection);
      messages = new Entities.Chat.Messages;
      API = {
        fetchMessages: function(projId) {
          var defer;
          defer = $.Deferred();
          messages.fetch({
            reset: true,
            data: {
              project_id: projId
            },
            success: function(collection, response, options) {
              return defer.resolve(messages);
            },
            error: function(collection, response, options) {
              return defer.resolve(null);
            }
          });
          return defer.promise();
        },
        sendMessage: function(objId, body) {
          var defer, message;
          defer = $.Deferred();
          message = new Entities.Chat.Message({
            project_id: objId,
            body: body
          });
          message.save(null, {
            success: function() {
              messages.add(message);
              return defer.resolve(message);
            },
            error: function() {
              return defer.resolve(null);
            }
          });
          return defer.promise();
        }
      };
      Devfeed.reqres.setHandler("chat:messages:fetch", function(objId) {
        return API.fetchMessages(objId);
      });
      return Devfeed.reqres.setHandler("chat:message:send", function(projId, body) {
        return API.sendMessage(projId, body);
      });
    });
    return Devfeed.Entities.Chat;
  });

}).call(this);
