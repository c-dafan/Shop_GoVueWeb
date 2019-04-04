import axios from 'axios'

let host = 'http://127.0.0.1:8088'

export const login = (params) => {
  return axios.post(`${host}/v1/login/`, params)
}

export const regist = (params) => {
  return axios.post(`${host}/v1/user/`, params)
}

export const good_list = () => {
  return axios.get(`${host}/v1/good`)
}

export const good_list_page = (page_id) => {
  return axios.get(`${host}/v1/good?offset=` + page_id)
}

export const good_list_v2 = (params) => {
  let page_id = params['page'] === undefined ? 0 : params['page']
  let sort_by = params['sort'] === undefined ? 'Id' : params['sort']
  let order = params['order'] === undefined ? 'asc' : params['order']
  let min_price = params['min_price'] == '' ? 0 : params['min_price']
  let max_price = params['max_price'] == '' ? '' : ',Price__lte:' + params['max_price']
  // Price__gte:10,Price__lte:
  return axios.get(`${host}/v1/good?offset=` + page_id + '&sortby=' + sort_by + '&order=' + order
    + '&query=Price__gte:' + min_price + max_price)
  // return axios.get(`${host}/v1/good?offset=` + page_id + '&sortby=' + sort_by + '&order=' + order)
}

export const good_detail = (gid) => {
  return axios.get(`${host}/v1/good/` + gid)
}

export const add_cart = (params) => {
  return axios.post(`${host}/v1/cart`, params)
}

export const get_cart = (uid) => {
  return axios.get(`${host}/v1/cart/` + uid)
}

export const delete_cart_cid = (cid) => {
  return axios.delete(`${host}/v1/cart/` + cid)
}
// export const put_cart_num= (params)=>{
//   return axios.put(`${host}/v1/cart`,params)
// }
export const patch_cart_num = (params) => {
  return axios.patch(`${host}/v1/cart`, params)
}

export const get_address_uid = (uid) => {
  return axios.get(`${host}/v1/address/uid/` + uid)
}

export const post_address = (params) => {
  return axios.post(`${host}/v1/address/`, params)
}

export const delete_address = (aid) => {
  return axios.delete(`${host}/v1/address/` + aid)
}

export const patch_user_address = (uid, aid) => {
  return axios.patch(`${host}/v1/user`, {Id: uid, Aid: {Id: aid}})
}

export const buy_now_post = (params) => {
  return axios.post(`${host}/v1/order`, params)
}

export const post_buy_item = (params) => {
  return axios.post(`${host}/v1/order_item`, params)
}

export const get_order_uid = (uid) => {
  return axios.get(`${host}/v1/order/uid/` + uid)
}

export const post_comment = (params) => {
  return axios.post(`${host}/v1/comment`, params)
}


