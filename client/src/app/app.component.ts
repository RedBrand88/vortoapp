import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface ITodoItem {
  text: string
  complete: Number
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public text = ''
  // sqlite does not have type bool just 1 and 0
  public complete = 0
  public todos: ITodoItem[] = []

  constructor(private httpClient: HttpClient) { }

  async ngOnInit() {
    const loaded = await this.loadTodos()
    console.log(loaded)
  }

  async loadTodos() {
    this.todos = await this.httpClient
      .get<ITodoItem[]>('/api/todo')
      .toPromise()
  }

  async addTodo() {
    this.httpClient.post('/api/todo', {
      title: this.text,
      complete: this.complete
    })
    await this.loadTodos()
    this.text = ''
    this.complete = 0
  }
}
