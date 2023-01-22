package examples.gets;

import com.intuit.karate.junit5.Karate;

class GetsRunner {
    
    @Karate.Test
    Karate testUsers() {
        return Karate.run("gets").relativeTo(getClass());
    }    

}
