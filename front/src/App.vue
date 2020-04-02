<template>
  <v-app id="inspire">
    <v-app-bar app color="indigo" dark>
      <v-toolbar-title>Application</v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container class="fill-height" fluid>
        <v-row align="center" justify="center">
          <v-col class="text-center">
            <v-simple-table>
              <template v-slot:default>
                <thead>
                  <tr>
                    <th>Name</th>
                    <th>Barcode</th>
                    <th>Price</th>
                    <th><v-btn color="success">新規作成</v-btn></th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in items" :key="item.id">
                    <td>{{ item.name }}</td>
                    <td>{{ item.barcode }}</td>
                    <td>{{ item.price }}</td>
                    <td>
                      <v-btn color="warning">編集</v-btn>
                      <v-btn color="error">削除</v-btn>
                    </td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
              <v-card>
                <v-text-field label="名称" v-model="item.name"/>
                <v-text-field label="バーコード" v-model="item.barcode"/>
                <v-text-field label="金額" v-model="item.price"/>
                <button @click="create">作成</button>
              </v-card>
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
      item:{}
    }
  },
  created(){
    this.read()
  },
  methods:{
    read(){
      fetch(new Request("/api/items",{
          method: "POST",
          dataType : "json",
      })).then(response=>{
          return response.json();
      }).then((items)=> {
        this.items = items
      })
    },
    create(){
      console.log(this.item)
      fetch(new Request("/api/items/create",{
          method: "POST",
          dataType : "json",
          body: JSON.stringify(this.item)
      })).then(response=>{
        return response.json();
      }).then(()=> {
        this.read()
      }) 
    }
  }
}
</script>

