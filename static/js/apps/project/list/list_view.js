(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "tpl!apps/project/list/templates/list.tpl", "tpl!apps/project/list/templates/empty.tpl", "tpl!apps/project/list/templates/project.tpl"], function(Devfeed, listTpl, emptyTpl, projectTpl) {
    Devfeed.module("ProjectApp.List.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref, _ref1, _ref2;
      View.Empty = (function(_super) {
        __extends(Empty, _super);

        function Empty() {
          _ref = Empty.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Empty.prototype.className = "empty";

        Empty.prototype.template = emptyTpl;

        return Empty;

      })(Marionette.ItemView);
      View.Project = (function(_super) {
        __extends(Project, _super);

        function Project() {
          _ref1 = Project.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        Project.prototype.tagName = "li";

        Project.prototype.template = projectTpl;

        Project.prototype.events = {
          "click a": "projectClicked"
        };

        Project.prototype.modelEvents = {
          "change": "render"
        };

        Project.prototype.projectClicked = function(e) {
          e.preventDefault();
          if (this.model.get("issynced")) {
            return Devfeed.trigger("project:show", this.model.get("id"));
          }
        };

        Project.prototype.onRender = function() {
          if (!this.model.get("issynced")) {
            return this.$(".syncing").html("&nbsp;").spin({
              lines: 8,
              length: 3,
              width: 3,
              radius: 4,
              color: "#9ea7b3",
              corners: 1
            });
          }
        };

        return Project;

      })(Marionette.ItemView);
      return View.List = (function(_super) {
        __extends(List, _super);

        function List() {
          _ref2 = List.__super__.constructor.apply(this, arguments);
          return _ref2;
        }

        List.prototype.id = "project-list";

        List.prototype.className = "row collapse";

        List.prototype.template = listTpl;

        List.prototype.emptyView = View.Empty;

        List.prototype.itemView = View.Project;

        List.prototype.itemViewContainer = ".projects";

        List.prototype.events = {
          "click .setup-msg a": "setupClicked"
        };

        List.prototype.setupClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("settings:general");
        };

        List.prototype.onCompositeRendered = function() {
          var userSession;
          userSession = Devfeed.request("user:session");
          if (!userSession.get("apitoken")) {
            this.$(".empty-msg").addClass("hide");
            return this.$(".setup-msg").removeClass("hide");
          }
        };

        return List;

      })(Marionette.CompositeView);
    });
    return Devfeed.ProjectApp.List.View;
  });

}).call(this);
