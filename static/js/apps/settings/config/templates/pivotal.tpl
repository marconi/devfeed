<div class="small-10 large-7 small-centered columns">
  <div class="row collapse">
    <div class="title small-12 columns">
      <h3><i class="icon-wrench"></i> Settings</h3>
    </div>
  </div>
  <div class="row collapse">
    <div class="small-12 columns">
      <div class="section-container vertical-tabs">
        <section class="general">
          <p class="title"><a href="#">General</a></p>
          <div class="content"></div>
        </section>
        <section class="active pivotal">
          <p class="title"><a href="#">Pivotal Tracker</a></p>
          <div class="content">
            <form method="post">
              <div class="alert-region"></div>
              <div class="small-12 columns">
                <input name="apitoken" type="text" placeholder="Pivotal Tracker API Token" value="<%= apitoken %>">
              </div>
              <div class="small-12 columns actions">
                <button type="submit" class="update">Update</button>
              </div>
            </form>
          </div>
        </section>
      </div>
    </div>
  </div>
</div>

