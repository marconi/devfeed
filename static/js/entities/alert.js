(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed"], function(Devfeed) {
    Devfeed.module("Entities", function(Entities, Devfeed, Backbone, Marionette, $, _) {
      var _ref;
      return Entities.Alert = (function(_super) {
        __extends(Alert, _super);

        function Alert() {
          _ref = Alert.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Alert.prototype.defaults = {
          type: null,
          message: null
        };

        return Alert;

      })(Backbone.Model);
    });
    return Devfeed.Entities.Alert;
  });

}).call(this);
