<template>
  <v-app id="inspire">
    <v-app-bar app color="indigo" dark>
      <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container class="fill-height" fluid>
        <v-row align="center" justify="center">
          <v-col class="text-center">
            <v-simple-table height="300px">
              <template v-slot:default>
                <thead>
                  <tr>
                    <th>Name</th>
                    <th>Barcode</th>
                    <th><v-btn color="success">新規作成</v-btn></th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in items" :key="item.id">
                    <td>{{ item.name }}</td>
                    <td>{{ item.barcode }}</td>
                    <td>
                      <v-btn color="warning">編集</v-btn>
                      <v-btn color="error">削除</v-btn>
                    </td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  data(){
    return {
      items:[],
    }
  },
  created(){
    var req = new Request("/api/items",{
        method: "GET",
        dataType : "json",
    });
    fetch(req).then(response=>{
        return response.json();
    }).then((items)=> {
      this.items = items
    })
  }
}
</script>

