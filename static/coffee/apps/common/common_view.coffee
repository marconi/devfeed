define [
  "devfeed",
  "alert",
  "tpl!apps/common/templates/alert.tpl",
  "tpl!apps/common/templates/preloader.tpl"
], (Devfeed, Alert, alertTpl, preloaderTpl) ->

  Devfeed.module "Common.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.Alert extends Marionette.ItemView
      className: "alert-box radius"
      template: alertTpl
      attributes:
        "data-alert": ""
      onRender: ->
        @$el.addClass(@model.get("type"))

    class View.FormViewMixin extends Marionette.Layout
      regions:
        alertRegion: ".alert-region"

      onFormDataInvalid: (result) ->
        @resetForm()

        # display any error message
        if result.message?
          alert = new Alert
            message: result.message.body
            type: result.message.type
          alertView = new View.Alert model: alert
          @alertRegion.show(alertView)
          delete result.message

        # display any field errors
        _.each result.errors, (error, field) ->
          $errorEl = $("<span>", {class: "error", text: error})
          @$("input[name=" + field + "]").addClass("error").after($errorEl)

      onFormDataValid: (result) ->
        @resetForm()
        @clearFields()

        if result.message?
          alert = new Alert
            message: result.message.body
            type: result.message.type
          alertView = new View.Alert model: alert
          @alertRegion.show(alertView)

      showPreloader: ->
        @ui.primaryBtn
          .attr("data-label", @ui.primaryBtn.html())
          .html("&nbsp;")
          .addClass("disabled")
          .spin("small")

      clearErrors: ->
        @$("input").removeClass("error")
        @$("span.error").remove()
        @alertRegion.close()

      resetForm: ->
        # remove preloader and restore button
        @ui.primaryBtn
          .spin(false)
          .removeClass("disabled")
          .html(@ui.primaryBtn.attr("data-label"))
          .removeAttr("data-label")

        # clear errors
        @clearErrors()

    class View.Preloader extends Marionette.ItemView
      id: "preloader"
      className: "row collapse"
      template: preloaderTpl
      onDomRefresh: ->
        @$(".loading").spin
           lines: 10
           length: 13
           width: 8
           radius: 13

  return Devfeed.Common.View
