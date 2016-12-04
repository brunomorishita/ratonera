package com.example.android.ratoneira;

import org.json.JSONException;
import org.json.JSONObject;

// Reference:
// https://www.javacodegeeks.com/2013/10/android-json-tutorial-create-and-parse-json-data.html

public class JsonUtil {

    public static String toJSon(SensorsData data) {
        try {
            // Here we convert Java Object to JSON
            JSONObject jsonObj = new JSONObject();
            jsonObj.put("id", data.getDeviceId());
            jsonObj.put("time", data.getTimestamp()); // Set the first name/pair

            // and finally we add the phone number
            // In this case we need a json array to hold the java list
            SensorsData.Accelerometer acc = data.getAcc();
            JSONObject jobj = new JSONObject();
            jobj.put("x", acc.getX());
            jobj.put("y", acc.getY());
            jobj.put("z", acc.getZ());
            jsonObj.put("accel", jobj);

            SensorsData.GPS gps = data.getGPS();
            jobj = new JSONObject();
            jobj.put("lat", gps.getLat());
            jobj.put("lgt", gps.getLgt());
            jsonObj.put("gps", jobj);



/*            jsonObj.put("surname", data.getSurname());

            JSONObject jsonAdd = new JSONObject(); // we need another object to store the address
            jsonAdd.put("address", data.getAddress().getAddress());
            jsonAdd.put("city", data.getAddress().getCity());
            jsonAdd.put("state", data.getAddress().getState());

            // We add the object to the main object
            jsonObj.put("address", jsonAdd);

            // and finally we add the phone number
            // In this case we need a json array to hold the java list
            JSONArray jsonArr = new JSONArray();

            for (SensorsData.Accelerometer acc : data.getAcc() ) {
                JSONObject pnObj = new JSONObject();
                pnObj.put("num", pn.getNumber());
                pnObj.put("type", pn.getType());
                jsonArr.put(pnObj);
            }

            jsonObj.put("phoneNumber", jsonArr);*/

            return jsonObj.toString();

        } catch (JSONException ex) {
            ex.printStackTrace();
        }

        return null;

    }
}