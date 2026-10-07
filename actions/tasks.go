package actions

import (
	"net/http"
	"strings"

	"todo/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

// TasksIndex muestra la lista de tareas y el formulario de creación.
func TasksIndex(c buffalo.Context) error {
	return renderIndex(c, http.StatusOK, &models.Task{}, validate.NewErrors(), nil, validate.NewErrors())
}

// TasksCreate crea una tarea nueva a partir del formulario.
func TasksCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	task := &models.Task{}
	if err := c.Bind(task); err != nil {
		return err
	}
	task.Title = strings.TrimSpace(task.Title)

	verrs, err := tx.ValidateAndCreate(task)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return renderIndex(c, http.StatusUnprocessableEntity, task, verrs, nil, validate.NewErrors())
	}

	c.Flash().Add("success", "Tarea creada.")
	return c.Redirect(http.StatusSeeOther, "/")
}

// TasksEdit muestra la lista con el título de una tarea en modo edición.
func TasksEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	task := &models.Task{}
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return renderIndex(c, http.StatusOK, &models.Task{}, validate.NewErrors(), task, validate.NewErrors())
}

// TasksUpdate cambia el título de una tarea.
func TasksUpdate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	task := &models.Task{}
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := c.Bind(task); err != nil {
		return err
	}
	task.Title = strings.TrimSpace(task.Title)

	verrs, err := tx.ValidateAndUpdate(task)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return renderIndex(c, http.StatusUnprocessableEntity, &models.Task{}, validate.NewErrors(), task, verrs)
	}

	c.Flash().Add("success", "Tarea actualizada.")
	return c.Redirect(http.StatusSeeOther, "/")
}

// renderIndex pinta la lista de tareas con el formulario de creación y,
// opcionalmente, una tarea en modo edición.
func renderIndex(c buffalo.Context, status int, task *models.Task, errors *validate.Errors, editTask *models.Task, editErrors *validate.Errors) error {
	tx := c.Value("tx").(*pop.Connection)

	tasks := models.Tasks{}
	if err := tx.Order("created_at asc, id asc").All(&tasks); err != nil {
		return err
	}

	editingID := ""
	editTitle := ""
	if editTask != nil {
		editingID = editTask.ID.String()
		editTitle = editTask.Title
	}

	c.Set("tasks", tasks)
	c.Set("task", task)
	c.Set("errors", errors)
	c.Set("editingID", editingID)
	c.Set("editTitle", editTitle)
	c.Set("editErrors", editErrors)
	return c.Render(status, r.HTML("tasks/index.plush.html"))
}

// TasksToggle alterna una tarea entre completada y pendiente.
func TasksToggle(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	task := &models.Task{}
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	task.Completed = !task.Completed
	if err := tx.Update(task); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

// TasksDestroy elimina una tarea.
func TasksDestroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	task := &models.Task{}
	if err := tx.Find(task, c.Param("task_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(task); err != nil {
		return err
	}

	c.Flash().Add("success", "Tarea eliminada.")
	return c.Redirect(http.StatusSeeOther, "/")
}
