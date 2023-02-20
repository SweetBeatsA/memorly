import { Component } from '@angular/core';




@Component({ templateUrl: 'login.component.html',
styleUrls: ["./login.component.css"] })
export class LoginComponent {
    public showPassword: boolean = false;

    public togglePasswordVisibility(): void {
        this.showPassword = !this.showPassword;
      }
}