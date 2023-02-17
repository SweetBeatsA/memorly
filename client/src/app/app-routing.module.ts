import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { HomeComponent } from './home';
import { LoginComponent } from './login';
import { RegisterComponent } from './register';
import { CreateCardComponent } from './create-card';


const routes: Routes = [
  { path: '', component: HomeComponent },
    { path: 'login', component: LoginComponent },
    { path: 'register', component: RegisterComponent },
    { path: 'create-card', component: CreateCardComponent },

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];
export const appRoutingModule = RouterModule.forRoot(routes);   //Is it correct to have both export lines in this file?

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
