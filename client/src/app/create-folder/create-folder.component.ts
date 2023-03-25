import { Component } from '@angular/core';
import axios from 'axios';
import {Router} from '@angular/router';

@Component({ 
  templateUrl: 'create-folder.component.html',
styleUrls: ["./create-folder.component.css"] })

export class CreateFolderComponent {


  constructor(private router: Router) {}


    createFolder(title : string) {

        const data = { title };

        // Define the headers with the access token
        const headers = { Authorization: sessionStorage.getItem('accessToken') };
        
        // Make the POST request
        axios.post('http://api.memorly.kro.kr/folder', data, { headers })
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
