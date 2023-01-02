import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { environment } from 'src/environments/environment';
import {Property} from "../models/property";

@Injectable({
  providedIn: 'root',
})
export class PropertiesService {

  constructor(private http: HttpClient) {}

  getAllProperties(uid:string) {
    const headers = new HttpHeaders()
    headers.set('Content-Type', 'application/json');
    const options = {headers: headers};
    return this.http.get<Property[]>(`${environment.apiURL}/users/${uid}/properties`, options);
  }

}
