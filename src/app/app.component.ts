import { Component, OnInit } from '@angular/core';
import {HelloWorldService} from './hello-world.service';
import {IArticle} from './article'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  
  title = 'angular-go-boilderplate';

  articles: IArticle[];

  constructor(private hw: HelloWorldService) {}
  
  ngOnInit() {
    
    console.log("Inside Init")
    this.hw.getTitle()
      .subscribe(data => this.articles = data)

    console.log(this.articles);
  }
  
}
