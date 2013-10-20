<a href="/projects/#/stories/<%= id %>" class="name"><%= name %></a>
<ul class="tasks hide">
  <% if (tasks.length > 0) { %>
    <% tasks.each(function(task) { %>
      <li>
        <% if (task.get("complete")) { %>
          <a href="#" class="task complete">
            <input type="checkbox" name="task" checked="true">
        <% } else { %>
          <a href="#" class="task">
            <input type="checkbox" name="task">
        <% } %>
          <%= task.get("description") %>
        </a>
      </li>
    <% }); %>
  <% } else { %>
    <li class="empty">This story doesn't have tasks.</li>
  <% } %>
</ul>