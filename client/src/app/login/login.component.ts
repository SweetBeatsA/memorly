import {MatSnackBar} from '@angular/material/snack-bar';
import { Component } from '@angular/core';
import {Router} from '@angular/router';
import axios from 'axios';

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

      

  constructor(private snackBar: MatSnackBar, 
    private router: Router, 
    private authService: AuthService) {}

  login(email1 : string,  password1 : string) {
    const data = {
      email: email1,
      password: password1
    };
    //const email = email1;
    //const password = password1;
    

    //this.authService.login(email, password)
    axios.post('http://api.memorly.kro.kr/users/login', data)
      .then((response) => {
        // Navigate to the home page or other protected routes
            console.log(response);
            console.log(response.data.data.accessToken);
            console.log(response.data.data.refreshToken);

            // need to track jwts
            sessionStorage.setItem('accessToken', response.data.data.accessToken);
            sessionStorage.setItem('refreshToken', response.data.data.refreshToken);

            let snackBarRef = this.snackBar.open('Login successful', 'x', {duration: 10000});

            axios.post('http://api.memorly.kro.kr/user', sessionStorage.getItem('accessToken'))
              .then((response2)=> {
                console.log(response2);
                console.log(response2.data.data.user.name);
                localStorage.setItem('username', response2.data.data.user.name);
                localStorage.setItem('isLoggedIn', 'true');
              })
            .catch((error) => {
              console.error(error);
              if(error.response2.status >= 400){
                let snackBarRef = this.snackBar.open('Error: ' + error.response2.message, 'x', {duration: 10000});
              } 
              //let snackBarRef = this.snackBar.open('Error getting user name', 'x', {duration: 10000});
            })
            this.router.navigateByUrl('library');
      })
      .catch((error) => {
        console.error(error);
        //let snackBarRef = this.snackBar.open('This one\'s on us... try again later', 'x', {duration: 10000});
        if(error.response.status >= 500){
          let snackBarRef = this.snackBar.open('This one\'s on us... try again later', 'x', {duration: 10000});
        }
        else if(error.response.status === 404){
          let snackBarRef = this.snackBar.open('Account with given email not found', 'x', {duration: 10000});
        }
        else if(error.response.status === 401){
          let snackBarRef = this.snackBar.open('Incorrect password', 'x', {duration: 10000});
        }
        else if(error.response.message === 'Binding Error'){
          let snackBarRef = this.snackBar.open('Please fill out each form', 'x', {duration: 10000});
        };
      });


  }
}



