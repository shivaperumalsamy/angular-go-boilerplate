import { Component, OnInit } from '@angular/core';
import {ArticleService} from './services/articles.service';
import {IArticle} from './article'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  
  title = 'angular-go-boilderplate';

  articles: IArticle[];

  constructor(private hw: ArticleService) {}

  ngOnInit() {
  
    this.hw.getTitle()
      .subscribe(data => this.articles = data)

    console.log(this.articles);
  }
  
}
