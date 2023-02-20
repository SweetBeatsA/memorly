import { Component } from '@angular/core';

@Component({ templateUrl: 'register.component.html',
styleUrls: ["./register.component.css"] })
export class RegisterComponent {

    public showPassword: boolean = false;

    public togglePasswordVisibility(): void {
        this.showPassword = !this.showPassword;
      }
}