define [
  "devfeed",
  "tpl!apps/project/show/templates/sidebar.tpl",
  "tpl!apps/project/show/templates/chatinfo.tpl",
  "tpl!apps/project/show/templates/chatbox.tpl",
  "tpl!apps/project/show/templates/show.tpl"
], (Devfeed, sidebarTpl, chatinfoTpl, chatboxTpl, showTpl) ->

  Devfeed.module "ProjectApp.Show.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.Sidebar extends Marionette.ItemView
      id: "sidebar"
      template: sidebarTpl
      events:
        "click #hide-sidebar": "hidesidebarClicked"

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
