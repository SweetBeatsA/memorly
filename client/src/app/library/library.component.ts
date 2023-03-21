import { Component } from '@angular/core';
import axios from 'axios';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';

@Component({ templateUrl: 'library.component.html',
styleUrls: ["./library.component.css"] })

export class LibraryComponent {

    createFolder() {


        const data = { param1: 'name', param2: 'description' };

        // Define the headers with the access token
        const headers = { Authorization: sessionStorage.getItem('accessToken') };
        
        // Make the POST request
        axios.post('http://api.memorly.kro.kr/users/library', data, { headers })
          .then(response => {
            // Request was successful, log the response data
            console.log(response.data);
          })
          .catch(error => {
            // Request failed, log the error message
            console.error(error.message);
          });



    }






}
