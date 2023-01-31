import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class HelloWorldService {

  constructor(private http: HttpClient) { }

  getTitle() {
    console.log("CALLED");
    return this.http.get(`${environment.serverUrl}/hello-world`);
  }
}
