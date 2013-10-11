define [
  "devfeed",
  "alert",
  "common_view",
  "tpl!apps/settings/config/templates/general.tpl",
  "tpl!apps/settings/config/templates/pivotal.tpl"
], (Devfeed, Alert, CommonView, generalTpl, pivotalTpl) ->

  Devfeed.module "SettingsApp.Config.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.General extends CommonView.FormViewMixin
      id: "general-settings"
      className: "row collapse"
      template: generalTpl
      events:
        "click .update": "updateClicked"
        "click .general .title a": "generalClicked"
        "click .pivotal .title a": "pivotalClicked"
      ui:
        passwordInput: "input[name=password]"
        primaryBtn: ".update"

      updateClicked: (e) ->
        e.preventDefault()
        if @ui.primaryBtn.hasClass("disabled")
          return

        @clearErrors()
        @showPreloader()

        data = Backbone.Syphon.serialize(@)
        @trigger("settings:general:update", data)

      generalClicked: (e) ->
        e.preventDefault()

      pivotalClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("settings:pivotal")

      clearFields: ->
        # we only clear password
        @ui.passwordInput.val("")

    class View.Pivotal extends CommonView.FormViewMixin
      id: "pivotal-settings"
      className: "row collapse"
      template: pivotalTpl
      events:
        "click .update": "updateClicked"
        "click .general .title a": "generalClicked"
        "click .pivotal .title a": "pivotalClicked"
      ui:
        primaryBtn: ".update"

      updateClicked: (e) ->
        e.preventDefault()
        if @ui.primaryBtn.hasClass("disabled")
          return

        @clearErrors()
        @showPreloader()

        data = Backbone.Syphon.serialize(@)
        @trigger("settings:pivotal:update", data)

      generalClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("settings:general")

      pivotalClicked: (e) ->
        e.preventDefault()

      clearFields: ->
        # we do nothing here since there's nothing to clear

  return Devfeed.SettingsApp.Config.View
