import React, { useState, useEffect } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Button, List, Segment, Container, Divider } from "semantic-ui-react";

const endpoint = "http://localhost:9000";

const ToDo = () => {
  const [task, setTask] = useState("");
  const [items, setItems] = useState([]);
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    getTasks();
  }, []);

  useEffect(() => {
    if (isSubmitting) {
      setIsSubmitting(false);
    }
  }, [items]);

  const onSubmit = (e) => {
    e.preventDefault();
    if (task && !isSubmitting) {
      setIsSubmitting(true);
      addTask(task);  // Pass task here
    }
};

  const getTasks = async () => {
    try {
      const response = await axios.get(`${endpoint}/api/task`);
      setItems(response.data);
    } catch (error) {
      console.error("Error fetching tasks:", error);
    }
  };

  const addTask = async (taskValue) => {
    try {
        console.log("Task value to be sent:", taskValue);
        await axios.post(`${endpoint}/api/tasks`, { task: taskValue });
        getTasks(); 
    } catch (error) {
        console.error("Error adding task:", error);
    }
};
  

  const completeTask = async (id) => {
    try {
      await axios.put(`${endpoint}/api/tasks/${id}`);
      getTasks();
    } catch (error) {
      console.error("Error completing task:", error);
    }
  };

  const undoTask = async (id) => {
    try {
      await axios.put(`${endpoint}/api/undoTask/${id}`);
      getTasks();
    } catch (error) {
      console.error("Error undoing task completion:", error);
    }
  };

  const deleteTask = async (id) => {
    try {
      await axios.delete(`${endpoint}/api/deleteTask/${id}`);
      getTasks();
    } catch (error) {
      console.error("Error deleting task:", error);
    }
  };

  const deleteAllTasks = async () => {
    try {
      await axios.delete(`${endpoint}/api/deleteAllTasks`);
      getTasks();
    } catch (error) {
      console.error("Error deleting all tasks:", error);
    }
  };

  return (
    <Container style={{ marginTop: '50px' }}>
      <Header as="h2" color="yellow" textAlign="center">
        To Do List
      </Header>
      <Divider />
      <Form onSubmit={onSubmit}>
        <Form.Group>
          <Form.Input
            width={12}
            type="text"
            name="task"
            onChange={(e) => setTask(e.target.value)}
            value={task}
            fluid
            placeholder="Create task"
            disabled={isSubmitting}
          />
          <Form.Button type="submit" width={4} color="blue" disabled={isSubmitting}>Add Task</Form.Button>
        </Form.Group>
      </Form>
      <Button color="red" onClick={deleteAllTasks} fluid style={{ marginBottom: '20px' }}>Delete All Tasks</Button>
      <Segment piled>
        <List divided relaxed>
          {items && items.map((item) => (
            <List.Item key={item._id}>
              <List.Content floated='right'>
                {!item.completed && <Button size="tiny" color="green" onClick={() => completeTask(item._id)}>Complete</Button>}
                {item.completed && <Button size="tiny" color="yellow" onClick={() => undoTask(item._id)}>Undo</Button>}
                <Button size="tiny" color="red" onClick={() => deleteTask(item._id)}>Delete</Button>
              </List.Content>
              <List.Icon name='tasks' size='large' verticalAlign='middle' />
              <List.Content style={{ textDecoration: item.completed ? 'line-through' : 'none', fontSize: '18px', fontWeight: item.completed ? '400' : 'bold' }}>
                {item.task}
              </List.Content>
            </List.Item>
          ))}
        </List>
      </Segment>
    </Container>
  );
};

export default ToDo;
