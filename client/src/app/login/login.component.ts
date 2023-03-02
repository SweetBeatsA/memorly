
import { Component } from '@angular/core';

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';


@Component({ templateUrl: 'login.component.html',
styleUrls: ["./login.component.css"],

selector: 'app-login',
template: `
  <form (submit)="login()">
    <label for="email">Email</label>
    <input type="email" name="email" [(ngModel)]="email">

    <label for="password">Password</label>
    <input type="password" name="password" [(ngModel)]="password">

    <button type="submit">Login</button>
  </form>
`,})
export class LoginComponent {
    public showPassword: boolean = false;


    public togglePasswordVisibility(): void {
        this.showPassword = !this.showPassword;
      }

      

  constructor(private authService: AuthService) {}

  login(email1 : string,  password1 : string) {

    const email = email1;
    const password = password1;

    this.authService.login(email, password)
      .then(() => {
        // Navigate to the home page or other protected routes
      });
  }
}



