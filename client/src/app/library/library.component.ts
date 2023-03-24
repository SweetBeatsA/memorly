import { Component, OnInit } from '@angular/core';
import axios from 'axios';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';

import { MatGridListModule } from '@angular/material/grid-list';

@Component({ 
  templateUrl: 'library.component.html',
styleUrls: ["./library.component.css"] })

export class LibraryComponent implements OnInit {

  folders: any[] = [];

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.http.get('http://api.memorly.kro.kr/library').subscribe((response: any) => {
      this.folders = response.folders.map((folder: any) => ({
        id: folder.id,
        name: folder.name,
        icon: folder.type === 'document' ? 'description' : 'folder'
      }));
    });
  }
    //This is in a different file now
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
