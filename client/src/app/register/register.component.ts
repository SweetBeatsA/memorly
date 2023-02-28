import { Component } from '@angular/core';
import axios from 'axios';

@Component({ templateUrl: 'register.component.html',

//selector: 'app-root',
//  template: 
//    <button (click)="signUp()">Sign Up</button>,

styleUrls: ["./register.component.css"] })


export class RegisterComponent {

    public showPassword: boolean = false;

    public togglePasswordVisibility(): void {
        this.showPassword = !this.showPassword;
      }

      public signUp(): void {
        const data = {
          email: 'tester@gmail.com',
          password: 'testerPassword',
          name: 'tester'
        };
    
        axios.post('http://api.memorly.kro.kr/users/signup', data)
          .then(response => console.log(response))
          .catch(error => console.error(error));
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
