import {Injectable} from "@angular/core";
import {HttpHeaders} from "@angular/common/http";

@Injectable({
  providedIn: 'root',
})
export class BaseService {

  /**
   * It retrieves the base headers to do a request.
   *
   * @return HttpHeaders with the following keys preset:
   *  - Content-Type = application/json
   * */
  getBaseHeaders() {
    const headers = new HttpHeaders()
    headers.set('Content-Type', 'application/json');
    return headers
  }
}
