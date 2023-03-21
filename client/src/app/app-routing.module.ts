import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { HomeComponent } from './home';
import { LoginComponent } from './login';
import { RegisterComponent } from './register';
import { CreateCardComponent } from './create-card';
import { BrowseComponent } from './browse';
import { LibraryComponent } from './library';

import { CreateFolderComponent } from './create-folder';
import { CardViewComponent } from './card-view';

const routes: Routes = [
  { path: '', component: HomeComponent },
    { path: 'login', component: LoginComponent },
    { path: 'register', component: RegisterComponent },
    { path: 'create-card', component: CreateCardComponent },
    { path: 'browse', component: BrowseComponent },
    { path: 'library', component: LibraryComponent },
    { path: 'card-view', component: CardViewComponent },
    { path: 'create-folder', component: CreateFolderComponent },

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];
export const appRoutingModule = RouterModule.forRoot(routes);   //Is it correct to have both export lines in this file?

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
