(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed"], function(Devfeed) {
    Devfeed.module("Common.Model", function(Model, Devfeed, Backbone, Marionette, $, _) {
      var _ref, _ref1;
      Model.BaseModel = (function(_super) {
        __extends(BaseModel, _super);

        function BaseModel() {
          _ref = BaseModel.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        BaseModel.prototype.parse = function(response, options) {
          if ((response.s != null) && (response.d != null) && response.s === 200) {
            return response.d;
          }
          return response;
        };

        return BaseModel;

      })(Backbone.Model);
      return Model.BaseCollection = (function(_super) {
        __extends(BaseCollection, _super);

        function BaseCollection() {
          _ref1 = BaseCollection.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        BaseCollection.prototype.comparator = function(model) {
          return model.get("id");
        };

        BaseCollection.prototype.parse = function(response, options) {
          if ((response.s != null) && (response.d != null) && response.s === 200) {
            return response.d;
          }
          return response;
        };

        return BaseCollection;

      })(Backbone.Collection);
    });
    return Devfeed.Common.Model;
  });

}).call(this);
