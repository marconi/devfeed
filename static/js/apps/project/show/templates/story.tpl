<a href="/projects/#/stories/<%= id %>" class="name"><%= name %></a>
<% if (tasks.length > 0) { %>
  <ul class="tasks hide">
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
  </ul>
<% } %>