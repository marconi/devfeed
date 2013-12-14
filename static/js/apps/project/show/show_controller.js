(function() {
  define(["devfeed", "common_view", "project_show_view"], function(Devfeed, CommonView, ProjectShowView) {
    Devfeed.module("ProjectApp.Show", function(Show, Devfeed, Backbone, Marionette, $, _) {
      var bindOnStoriesRendered, renderStories;
      renderStories = function(sidebarView, filters, project) {
        var storiesView;
        storiesView = new ProjectShowView.Stories({
          collection: project.get("stories")
        });
        storiesView.on("stories:more", function() {
          var fetchingStories;
          fetchingStories = Devfeed.request("project:stories", project.get("id"), filters, false, false);
          return $.when(fetchingStories).done(function() {
            return storiesView.triggerMethod("more:stories");
          });
        });
        sidebarView.storiesRegion.show(storiesView);
        return storiesView;
      };
      bindOnStoriesRendered = function(storiesView, findStoryView) {
        return storiesView.on("stories:more:rendered", function() {
          return findStoryView.bindLiveFilter();
        });
      };
      return Show.Controller = {
        showProject: function(id) {
          var fetchingProject, preloaderView;
          preloaderView = new CommonView.Preloader();
          Devfeed.contentRegion.show(preloaderView);
          fetchingProject = Devfeed.request("project:entity", id);
          return $.when(fetchingProject).done(function(project) {
            var chatinfoView, fetchingMessages, findStoryView, messagesPreloaderView, projectShowView, sidebarView, storiesView;
            sidebarView = new ProjectShowView.Sidebar({
              model: project
            });
            chatinfoView = new ProjectShowView.Chatinfo;
            messagesPreloaderView = new CommonView.Preloader({
              message: "Loading messages...",
              innerClassName: "small-10 large-6"
            });
            projectShowView = new ProjectShowView.Show({
              model: project
            });
            Devfeed.contentRegion.show(projectShowView);
            projectShowView.sidebarRegion.show(sidebarView);
            projectShowView.chatinfoRegion.show(chatinfoView);
            projectShowView.chatboxRegion.show(messagesPreloaderView);
            projectShowView.on("message:send", function(body) {
              var sending;
              sending = Devfeed.request("chat:message:send", project.get("oid"), body);
              return $.when(sending).done(function(message) {
                return console.log("scrolling...");
              });
            });
            fetchingMessages = Devfeed.request("chat:messages:fetch", project.get("id"));
            $.when(fetchingMessages).done(function(messages) {
              var chatboxView;
              chatboxView = new ProjectShowView.Chatbox({
                collection: messages
              });
              projectShowView.chatboxRegion.show(chatboxView);
              return projectShowView.triggerMethod("enable:input");
            });
            storiesView = renderStories(sidebarView, [], project);
            findStoryView = new ProjectShowView.FindStory;
            bindOnStoriesRendered(storiesView, findStoryView);
            findStoryView.on("settings:shown", function() {
              return sidebarView.triggerMethod("settings:shown");
            });
            findStoryView.on("settings:hidden", function() {
              return sidebarView.triggerMethod("settings:hidden");
            });
            findStoryView.on("filters:changed", function(filters) {
              var fetchingStories, filterPreloaderView;
              filterPreloaderView = new ProjectShowView.FilterPreloader;
              sidebarView.storiesRegion.show(filterPreloaderView);
              project.get("stories").reset();
              fetchingStories = Devfeed.request("project:stories", project.get("id"), filters, true, true);
              return $.when(fetchingStories).done(function() {
                storiesView = renderStories(sidebarView, findStoryView.filters, project);
                bindOnStoriesRendered(storiesView, findStoryView);
                if (!$("#find-story .settings").hasClass("hide")) {
                  return storiesView.$el.addClass("settings-shown");
                }
              });
            });
            sidebarView.findStoryRegion.show(findStoryView);
            return Devfeed.execute("ws:project:subscribe", id);
          });
        }
      };
    });
    return Devfeed.ProjectApp.Show.Controller;
  });

}).call(this);
