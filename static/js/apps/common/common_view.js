(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "common_utils", "alert", "tpl!apps/common/templates/alert.tpl", "tpl!apps/common/templates/preloader.tpl"], function(Devfeed, CommonUtils, Alert, alertTpl, preloaderTpl) {
    Devfeed.module("Common.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref, _ref1, _ref2;
      View.Alert = (function(_super) {
        __extends(Alert, _super);

        function Alert() {
          _ref = Alert.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Alert.prototype.className = "alert-box radius";

        Alert.prototype.template = alertTpl;

        Alert.prototype.attributes = {
          "data-alert": ""
        };

        Alert.prototype.onRender = function() {
          return this.$el.addClass(this.model.get("type"));
        };

        return Alert;

      })(Marionette.ItemView);
      View.FormViewMixin = (function(_super) {
        __extends(FormViewMixin, _super);

        function FormViewMixin() {
          _ref1 = FormViewMixin.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        FormViewMixin.prototype.regions = {
          alertRegion: ".alert-region"
        };

        FormViewMixin.prototype.onFormDataInvalid = function(result) {
          var alert, alertView;
          this.resetForm();
          if (result.message != null) {
            alert = new Alert({
              message: result.message.body,
              type: result.message.type
            });
            alertView = new View.Alert({
              model: alert
            });
            this.alertRegion.show(alertView);
            delete result.message;
          }
          return _.each(result.errors, function(error, field) {
            var $errorEl;
            $errorEl = $("<span>", {
              "class": "error",
              text: error
            });
            return this.$("input[name=" + field + "]").addClass("error").after($errorEl);
          });
        };

        FormViewMixin.prototype.onFormDataValid = function(result) {
          var alert, alertView;
          this.resetForm();
          this.clearFields();
          if (result.message != null) {
            alert = new Alert({
              message: result.message.body,
              type: result.message.type
            });
            alertView = new View.Alert({
              model: alert
            });
            return this.alertRegion.show(alertView);
          }
        };

        FormViewMixin.prototype.showPreloader = function() {
          return this.ui.primaryBtn.attr("data-label", this.ui.primaryBtn.html()).html("&nbsp;").addClass("disabled").spin(CommonUtils.SmallSpin);
        };

        FormViewMixin.prototype.clearErrors = function() {
          this.$("input").removeClass("error");
          this.$("span.error").remove();
          return this.alertRegion.close();
        };

        FormViewMixin.prototype.resetForm = function() {
          this.ui.primaryBtn.spin(false).removeClass("disabled").html(this.ui.primaryBtn.attr("data-label")).removeAttr("data-label");
          return this.clearErrors();
        };

        return FormViewMixin;

      })(Marionette.Layout);
      return View.Preloader = (function(_super) {
        __extends(Preloader, _super);

        function Preloader() {
          _ref2 = Preloader.__super__.constructor.apply(this, arguments);
          return _ref2;
        }

        Preloader.prototype.id = "preloader";

        Preloader.prototype.className = "row collapse";

        Preloader.prototype.template = preloaderTpl;

        Preloader.prototype.ui = {
          innerColumns: ".columns",
          message: ".message"
        };

        Preloader.prototype.onRender = function() {
          var innerClass, message;
          message = this.options.message || "Loading...";
          innerClass = this.options.innerClassName || "small-7 large-3";
          this.ui.innerColumns.addClass(innerClass);
          return this.ui.message.html(message);
        };

        Preloader.prototype.onDomRefresh = function() {
          return this.$(".loading").spin({
            lines: 10,
            length: 13,
            width: 8,
            radius: 13,
            corners: 1
          });
        };

        return Preloader;

      })(Marionette.ItemView);
    });
    return Devfeed.Common.View;
  });

}).call(this);
