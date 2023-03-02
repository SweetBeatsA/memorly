
import { Component } from '@angular/core';
import {MatSnackBar} from '@angular/material/snack-bar';
import {Router} from '@angular/router';
import axios from 'axios';

@Component({ templateUrl: 'register.component.html',

//selector: 'app-root',
//  template: 
//    <button (click)="signUp()">Sign Up</button>,

styleUrls: ["./register.component.css"] })


export class RegisterComponent {
    username: string = '';
    password: string = '';
    email: string = '';
    public isUsernameValid: boolean = true;
    public isEmailValid: boolean = true;
    public isPasswordValid: boolean = true;
    public showPassword: boolean = false;
    

    constructor(
      private snackBar: MatSnackBar,
      private router: Router,) {}

    public togglePasswordVisibility(): void {
        this.showPassword = !this.showPassword;
      }

      onKey(event: any, type: string){
        if(type === 'username'){
          this.username = event.target.value;
          this.validateUsername();
        }
        else if (type === 'password'){
          this.password = event.target.value;
          this.validatePassword();
        }
        else if(type === 'email'){
          this.email = event.target.value;
        }
      }

      validateUsername(): void{
        const pattern = new RegExp(/^[\w-.]*$/);
        if(pattern.test(this.username)){
          this.isUsernameValid = true;
        }
        else{
          this.isUsernameValid = false;
        }
      }

      validatePassword(): void{
        if(this.password.length < 8){
          this.isPasswordValid = false;
        }
        else{
          this.isPasswordValid = true;
        }
      }

      /*validateEmail(): void{
        const pattern = RegExp(/^[\w-.]*$/);
        if(pattern.test(this.email)){
          this.isEmailValid = true;
        }
        else{
          this.isEmailValid= false;
        }
      }
      */
      



      public signUp(email1 : string,  password1 : string, name1 : string): void {
        const data = {
          email: email1,
          password: password1,
          name: name1
        };
        if (!this.isUsernameValid || !this.isEmailValid || !this.isPasswordValid){}
        else if(this.email === "" || this.username === ""){
          let snackBarRef = this.snackBar.open('Please fill out all forms before submitting!', 'x', {duration: 10000});
        }
       else{
        
    
        axios.post('http://api.memorly.kro.kr/users/signup', data)
          .then((response) => {
            console.log(response);
            console.log(response.data.data.accessToken);
            console.log(response.data.data.refreshToken);

            // need to track jwt
            sessionStorage.setItem('accessToken', response.data.data.accessToken);
            sessionStorage.setItem('refreshToken', response.data.data.refreshToken);

            let snackBarRef = this.snackBar.open('Account creation successful', 'x', {duration: 10000});

            this.router.navigateByUrl('dashboard');

          })
          .catch((error) => {
            console.error(error);

            if(error.response.status >= 500){
              let snackBarRef = this.snackBar.open('This one\'s on us... try again later', 'x', {duration: 10000});
            }
            else if(error.response.status === 400){
              let snackBarRef = this.snackBar.open('Email has already been taken', 'x', {duration: 10000});
            }
            //else if(error.response.message === 'Binding Error'){
            //  let snackBarRef = this.snackBar.open('Please fill out each form', 'x', {duration: 10000});
            //}
          });
        }
      }

      

}

/*
      axios.post('api.memorly.kro.kr/users/signup', {
        email: 'tester@gmail.com',
        password: 'testerPassword',
        name: 'tester'
      }, {
        headers: {
          Authorization: 'your_token_here'
        }
      })
      .then(response => {
        console.log(response);
      })
      .catch(error => {
        console.error(error);
      });
      */