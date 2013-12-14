define ["devfeed"], (Devfeed) ->

  Devfeed.module "Entities", (Entities, Devfeed, Backbone, Marionette, $, _) ->

    class Entities.Alert extends Backbone.Model
      defaults:
        type: null
        message: null

  return Devfeed.Entities.Alert
