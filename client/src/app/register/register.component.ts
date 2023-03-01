import { Component } from '@angular/core';
import {MatSnackBar} from '@angular/material/snack-bar';
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

    public showPassword: boolean = false;
    

    constructor(private snackBar: MatSnackBar) {}

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
      



      public signUp(email1 : string,  password1 : string, name1 : string,): void {
        const data = {
          email: email1,
          password: password1,
          name: name1
        };
    
        axios.post('http://api.memorly.kro.kr/users/signup', data)
          .then((response) => {
              console.log(response);
          })
          .catch((error) => {
            console.error(error);
            let snackBarRef = this.snackBar.open('Error on sign up.  Please try again', 'x');
          });
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
}
