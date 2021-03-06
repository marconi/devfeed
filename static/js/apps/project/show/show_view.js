(function() {
  var __hasProp = {}.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

  define(["devfeed", "jquery_livefilter", "common_utils", "common_view", "tpl!apps/project/show/templates/sidebar.tpl", "tpl!apps/project/show/templates/findstory.tpl", "tpl!apps/project/show/templates/filterpreloader.tpl", "tpl!apps/project/show/templates/stories.tpl", "tpl!apps/project/show/templates/story.tpl", "tpl!apps/project/show/templates/empty.tpl", "tpl!apps/project/show/templates/chatinfo.tpl", "tpl!apps/project/show/templates/chatbox.tpl", "tpl!apps/project/show/templates/message.tpl", "tpl!apps/project/show/templates/show.tpl"], function(Devfeed, LiveFilter, CommonUtils, CommonView, sidebarTpl, findStoryTpl, filterPreloaderTpl, storiesTpl, storyTpl, emptyTpl, chatinfoTpl, chatboxTpl, messageTpl, showTpl) {
    Devfeed.module("ProjectApp.Show.View", function(View, Devfeed, Backbone, Marionette, $, _) {
      var _ref, _ref1, _ref2, _ref3, _ref4, _ref5, _ref6, _ref7, _ref8, _ref9;
      View.Empty = (function(_super) {
        __extends(Empty, _super);

        function Empty() {
          _ref = Empty.__super__.constructor.apply(this, arguments);
          return _ref;
        }

        Empty.prototype.className = "empty";

        Empty.prototype.tagName = "li";

        Empty.prototype.template = emptyTpl;

        return Empty;

      })(Marionette.ItemView);
      View.Story = (function(_super) {
        __extends(Story, _super);

        function Story() {
          _ref1 = Story.__super__.constructor.apply(this, arguments);
          return _ref1;
        }

        Story.prototype.tagName = "li";

        Story.prototype.template = storyTpl;

        Story.prototype.events = {
          "click .name": "nameClicked",
          "click .task": "taskClicked"
        };

        Story.prototype.onRender = function() {
          return this.$el.addClass(this.model.get("current_state"));
        };

        Story.prototype.nameClicked = function(e) {
          e.preventDefault();
          this.$el.toggleClass("open");
          return this.$(".tasks").toggleClass("hide");
        };

        Story.prototype.taskClicked = function(e) {
          var $a, $checkbox;
          e.preventDefault();
          $a = $(e.currentTarget);
          $a.toggleClass("complete");
          $checkbox = $a.find("input[type=checkbox]");
          return $checkbox.prop("checked", function(i, value) {
            return !value;
          });
        };

        return Story;

      })(Marionette.ItemView);
      View.Stories = (function(_super) {
        __extends(Stories, _super);

        function Stories() {
          _ref2 = Stories.__super__.constructor.apply(this, arguments);
          return _ref2;
        }

        Stories.prototype.id = "stories";

        Stories.prototype.template = storiesTpl;

        Stories.prototype.emptyView = View.Empty;

        Stories.prototype.itemView = View.Story;

        Stories.prototype.itemViewContainer = ".inner ul";

        Stories.prototype.events = {
          "click .more": "moreClicked"
        };

        Stories.prototype.moreClicked = function(e) {
          e.preventDefault();
          this.$(".more span").addClass("hide");
          this.$(".more div").spin(CommonUtils.SmallSpin);
          return this.trigger("stories:more");
        };

        Stories.prototype.onMoreStories = function() {
          this.$(".more div").spin(false);
          this.$(".more span").removeClass("hide");
          return this.trigger("stories:more:rendered");
        };

        Stories.prototype.onCompositeCollectionRendered = function() {
          if (this.collection.size() > 0) {
            return this.$(".more").removeClass("hide");
          } else {
            return this.$(".more").addClass("hide");
          }
        };

        return Stories;

      })(Marionette.CompositeView);
      View.FilterPreloader = (function(_super) {
        __extends(FilterPreloader, _super);

        function FilterPreloader() {
          _ref3 = FilterPreloader.__super__.constructor.apply(this, arguments);
          return _ref3;
        }

        FilterPreloader.prototype.id = "filter-preloader";

        FilterPreloader.prototype.template = filterPreloaderTpl;

        FilterPreloader.prototype.onDomRefresh = function() {
          this.$(".loading").spin(CommonUtils.SmallSpin);
          if (!$("#find-story .settings").hasClass("hide")) {
            return this.$el.addClass("settings-shown");
          }
        };

        return FilterPreloader;

      })(Marionette.ItemView);
      View.FindStory = (function(_super) {
        __extends(FindStory, _super);

        function FindStory() {
          _ref4 = FindStory.__super__.constructor.apply(this, arguments);
          return _ref4;
        }

        FindStory.prototype.id = "find-story";

        FindStory.prototype.className = "small-12 columns";

        FindStory.prototype.template = findStoryTpl;

        FindStory.prototype.events = {
          "click .settings-cog a": "settingsClicked",
          "change .settings input[type=checkbox]": "settingsChanged"
        };

        FindStory.prototype.filters = [];

        FindStory.prototype.onRender = function() {
          return this.bindLiveFilter();
        };

        FindStory.prototype.bindLiveFilter = function() {
          return this.$('.keyword input').fastLiveFilter('#stories .inner ul:first-child');
        };

        FindStory.prototype.settingsClicked = function(e) {
          e.preventDefault();
          if (this.$(".settings").hasClass("hide")) {
            this.trigger("settings:shown");
          } else {
            this.trigger("settings:hidden");
          }
          return this.$(".settings").toggleClass("hide");
        };

        FindStory.prototype.settingsChanged = function(e) {
          var filters;
          filters = _.filter(this.$(".settings input[type=checkbox]"), function(checkbox) {
            return $(checkbox).is(":checked");
          });
          filters = _.map(filters, function(filter) {
            return $(filter).attr("name");
          });
          this.filters = filters;
          return this.trigger("filters:changed", this.filters);
        };

        return FindStory;

      })(Marionette.ItemView);
      View.Sidebar = (function(_super) {
        __extends(Sidebar, _super);

        function Sidebar() {
          _ref5 = Sidebar.__super__.constructor.apply(this, arguments);
          return _ref5;
        }

        Sidebar.prototype.id = "sidebar";

        Sidebar.prototype.template = sidebarTpl;

        Sidebar.prototype.events = {
          "click #hide-sidebar": "hidesidebarClicked"
        };

        Sidebar.prototype.regions = {
          findStoryRegion: "#find-story-region",
          storiesRegion: "#stories-region"
        };

        Sidebar.prototype.filterPreloaderView = null;

        Sidebar.prototype.hidesidebarClicked = function(e) {
          e.preventDefault();
          return console.log("hiding...");
        };

        Sidebar.prototype.onSettingsShown = function() {
          this.storiesRegion.currentView.$el.addClass("settings-shown");
          if (this.filterPreloaderView) {
            return this.filterPreloaderView.$el.addClass("settings-shown");
          }
        };

        Sidebar.prototype.onSettingsHidden = function() {
          this.storiesRegion.currentView.$el.removeClass("settings-shown");
          if (this.filterPreloaderView) {
            return this.filterPreloaderView.$el.removeClass("settings-shown");
          }
        };

        return Sidebar;

      })(Marionette.Layout);
      View.Chatinfo = (function(_super) {
        __extends(Chatinfo, _super);

        function Chatinfo() {
          _ref6 = Chatinfo.__super__.constructor.apply(this, arguments);
          return _ref6;
        }

        Chatinfo.prototype.id = "chatinfo";

        Chatinfo.prototype.template = chatinfoTpl;

        Chatinfo.prototype.events = {
          "click #back-projects a": "backprojectsClicked"
        };

        Chatinfo.prototype.backprojectsClicked = function(e) {
          e.preventDefault();
          return Devfeed.trigger("projects:list");
        };

        return Chatinfo;

      })(Marionette.ItemView);
      View.Message = (function(_super) {
        __extends(Message, _super);

        function Message() {
          _ref7 = Message.__super__.constructor.apply(this, arguments);
          return _ref7;
        }

        Message.prototype.className = "message small-12 columns";

        Message.prototype.template = messageTpl;

        return Message;

      })(Marionette.ItemView);
      View.Chatbox = (function(_super) {
        __extends(Chatbox, _super);

        function Chatbox() {
          _ref8 = Chatbox.__super__.constructor.apply(this, arguments);
          return _ref8;
        }

        Chatbox.prototype.id = "chatbox";

        Chatbox.prototype.template = chatboxTpl;

        Chatbox.prototype.itemView = View.Message;

        Chatbox.prototype.itemViewContainer = "#messages";

        return Chatbox;

      })(Marionette.CompositeView);
      return View.Show = (function(_super) {
        __extends(Show, _super);

        function Show() {
          _ref9 = Show.__super__.constructor.apply(this, arguments);
          return _ref9;
        }

        Show.prototype.id = "project-details";

        Show.prototype.className = "row collapse";

        Show.prototype.template = showTpl;

        Show.prototype.regions = {
          sidebarRegion: "#sidebar-region",
          chatinfoRegion: "#chatinfo-region",
          chatboxRegion: "#chatbox-region"
        };

        Show.prototype.events = {
          "keypress #chatinput input": "keypressedInput"
        };

        Show.prototype.keypressedInput = function(e) {
          var $input, message;
          if (e.keyCode === 13) {
            e.preventDefault();
            $input = $(e.currentTarget);
            message = $input.val();
            $input.val("").focus();
            return this.trigger("message:send", message);
          }
        };

        Show.prototype.onEnableInput = function() {
          return this.$("#chatinput input").removeClass("disable").removeAttr("disabled");
        };

        return Show;

      })(Marionette.Layout);
    });
    return Devfeed.ProjectApp.Show.View;
  });

}).call(this);
