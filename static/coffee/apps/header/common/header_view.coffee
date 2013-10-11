define [
  "devfeed",
  "tpl!apps/header/common/templates/header.tpl",
  "tpl!apps/header/common/templates/userinfo.tpl"
], (Devfeed, headerTpl, userinfoTpl) ->

  Devfeed.module "HeaderApp.Common.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.Header extends Marionette.Layout
      id: "header"
      className: "small-12 columns"
      template: headerTpl
      events:
        "click #logo": "logoClicked"
      regions:
        userinfoRegion: "#userinfo-region"

      logoClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("home")

    class View.UserinfoView extends Marionette.ItemView
      id: "userinfo"
      template: userinfoTpl
      events:
        "click .logout a": "logoutClicked"
        "click .settings a": "settingsClicked"

      settingsClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("settings:general")
        @$("#user-name").click()

      logoutClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("auth:logout")

  return Devfeed.HeaderApp.Common.View
