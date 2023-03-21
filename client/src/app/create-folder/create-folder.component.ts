import { Component } from '@angular/core';
import axios from 'axios';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';

import {Router} from '@angular/router';

@Component({ 
  templateUrl: 'create-folder.component.html',
styleUrls: ["./create-folder.component.css"] })

export class CreateFolderComponent {


  constructor(private router: Router) {}


    createFolder(name1 : string,  description1 : string) {

        const data = { 
          name: name1, 
          description: description1 };

        // Define the headers with the access token
        const headers = { Authorization: sessionStorage.getItem('accessToken') };
        
        // Make the POST request
        axios.post('http://api.memorly.kro.kr/users/library', data, { headers })
          .then(response => {
            // Request was successful, log the response data
            console.log(response.data);

            this.router.navigateByUrl('library');

          })
          .catch(error => {
            // Request failed, log the error message
            console.error(error.message);
          });



    }






}
