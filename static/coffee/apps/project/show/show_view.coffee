define [
  "devfeed",
  "tpl!apps/project/show/templates/sidebar.tpl",
  "tpl!apps/project/show/templates/stories.tpl",
  "tpl!apps/project/show/templates/story.tpl",
  "tpl!apps/project/show/templates/empty.tpl",
  "tpl!apps/project/show/templates/chatinfo.tpl",
  "tpl!apps/project/show/templates/chatbox.tpl",
  "tpl!apps/project/show/templates/show.tpl"
], (Devfeed, sidebarTpl, storiesTpl, storyTpl, emptyTpl, chatinfoTpl, chatboxTpl, showTpl) ->

  Devfeed.module "ProjectApp.Show.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.Empty extends Marionette.ItemView
      className: "empty"
      tagName: "li"
      template: emptyTpl

    class View.Story extends Marionette.ItemView
      tagName: "li"
      template: storyTpl
      events:
        "click .name": "nameClicked"
        "click .task": "taskClicked"

      nameClicked: (e) ->
        e.preventDefault()
        @$el.toggleClass("open")
        @$(".tasks").toggleClass("hide")

      taskClicked: (e) ->
        e.preventDefault()
        aEl = $(e.currentTarget)
        aEl.toggleClass("complete")

        checkBox = aEl.find("input[type=checkbox]")
        checkBox.prop "checked", (i, value) ->
          return not value

    class View.Stories extends Marionette.CompositeView
      id: "stories"
      template: storiesTpl
      emptyView: View.Empty
      itemView: View.Story
      itemViewContainer: ".inner ul"

    class View.Sidebar extends Marionette.Layout
      id: "sidebar"
      template: sidebarTpl
      events:
        "click #hide-sidebar": "hidesidebarClicked"
      regions:
        findStoryRegion: "#find-story-region"
        storiesRegion: "#stories-region"

      hidesidebarClicked: (e) ->
        e.preventDefault()
        console.log "hiding..."

    class View.Chatinfo extends Marionette.ItemView
      id: "chatinfo"
      template: chatinfoTpl
      events:
        "click #back-projects a": "backprojectsClicked"

      backprojectsClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("projects:list")

    class View.Chatbox extends Marionette.ItemView
      id: "chatbox"
      template: chatboxTpl

    class View.Show extends Marionette.Layout
      id: "project-details"
      className: "row collapse"
      template: showTpl
      regions:
        sidebarRegion: "#sidebar-region"
        chatinfoRegion: "#chatinfo-region"
        chatboxRegion: "#chatbox-region"

  return Devfeed.ProjectApp.Show.View
