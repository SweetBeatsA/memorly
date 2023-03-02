import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private readonly apiUrl = 'http://api.memorly.kro.kr/users/login';

  constructor(private http: HttpClient) {}

  login(email: string, password: string) {
    return this.http.post<any>(this.apiUrl, { email, password })
      .toPromise()
      .then(response => {
        // Save the JWT token to local storage
        localStorage.setItem('token', response.token);
      });
  }

  logout() {
    // Remove the JWT token from local storage
    localStorage.removeItem('token');
  }

  isLoggedIn() {
    // Check if the JWT token is present in local storage
    return localStorage.getItem('token') !== null;
  }

}