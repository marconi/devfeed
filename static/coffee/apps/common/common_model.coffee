define ["devfeed"], (Devfeed) ->

  Devfeed.module "Common.Model", (Model, Devfeed, Backbone, Marionette, $, _) ->

    class Model.BaseModel extends Backbone.Model
      parse: (response, options) ->
        if response.s == 200
          return response.d

    class Model.BaseCollection extends Backbone.Collection
      comparator: (model) ->
        return model.get("id")
      parse: (response, options) ->
        if response.s == 200
          return response.d

  return Devfeed.Common.Model
