import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface ITodoItem {
  id: Number
  text: string
  complete: Number
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public id: number
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
    await this.httpClient.post('/api/todo', {
      text: this.text,
      complete: this.complete
    }).toPromise()

    await this.loadTodos()
    this.text = ''
    this.complete = 0
  }

  async updateTodos(event) {
    if (event.target.checked) {
      this.complete = 1
    } else {
      this.complete = 0
    }

    await this.httpClient.put('/api/todo', {
      id: event.target.id,
      complete: this.complete
    }).toPromise()
    this.complete = 0
  }

  logger = val => console.log(val.target.id)


}
