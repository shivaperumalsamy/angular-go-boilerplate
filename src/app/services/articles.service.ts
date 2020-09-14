import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from '../../environments/environment';
import { Observable } from 'rxjs';
import {IArticle} from './../article';

interface Article {
  title: string;
  desc:string;
  content:string;
}

@Injectable()
export class ArticleService {

  private _url: string = "http://localhost:4201/api/v1/article"

  constructor(private http: HttpClient) { }

  getTitle(): Observable<IArticle[]> {
    return this.http.get<IArticle[]>(this._url)
  }

}
