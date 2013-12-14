(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "tpl!apps/header/common/templates/header.tpl", "tpl!apps/header/common/templates/userinfo.tpl"], function(Devfeed, headerTpl, userinfoTpl) {
    Devfeed.module("HeaderApp.Common.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref, _ref1;
      View.Header = (function(_super) {
        __extends(Header, _super);

        function Header() {
          _ref = Header.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Header.prototype.id = "header";

        Header.prototype.className = "small-12 columns";

        Header.prototype.template = headerTpl;

        Header.prototype.events = {
          "click #logo": "logoClicked"
        };

        Header.prototype.regions = {
          userinfoRegion: "#userinfo-region"
        };

        Header.prototype.logoClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("home");
        };

        return Header;

      })(Marionette.Layout);
      return View.UserinfoView = (function(_super) {
        __extends(UserinfoView, _super);

        function UserinfoView() {
          _ref1 = UserinfoView.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        UserinfoView.prototype.id = "userinfo";

        UserinfoView.prototype.template = userinfoTpl;

        UserinfoView.prototype.events = {
          "click .settings a": "settingsClicked"
        };

        UserinfoView.prototype.settingsClicked = function(e) {
          e.preventDefault();
          Devfeed.trigger("settings:general");
          return this.$("#user-name").click();
        };

        return UserinfoView;

      })(Marionette.ItemView);
    });
    return Devfeed.HeaderApp.Common.View;
  });

}).call(this);
