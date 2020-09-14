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

  // articles: IArticle[] = [
  //   {title: 'title1', desc: 'Sentence 1', content: 'content1'},
  //   {title: 'title2', desc: 'Sentence 2', content: 'content2'},
  //   {title: 'title3', desc: 'Sentence 3', content: 'content3'},
  // ];

  

  ngOnInit() {
    
    this.articles = [
      {title: 'title1', desc: 'Sentence 1', content: 'content1'},
      {title: 'title2', desc: 'Sentence 2', content: 'content2'},
      {title: 'title3', desc: 'Sentence 3', content: 'content3'},
    ];

    console.log("Inside Init")
    this.hw.getTitle()
      .subscribe(data => this.articles = data)

    console.log(this.articles);
  }
  
}
