use rand::distributions::Alphanumeric;
use rand::Rng;
use std::ffi::CString;
use std::os::raw::c_char;

#[no_mangle]
pub extern "C" fn generate_short_url()-> *mut c_char{
    let short_url: String= rand::thread_rng()
        .sample_iter(&Alphanumeric)
        .take(7)
        .map(char::from)
        .collect();

    CString::new(short_url).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn free_string(s: *mut c_char){
    unsafe{
        if s.is_null(){
            return;
        }
        let _ = CString::from_raw(s);
    }
}
