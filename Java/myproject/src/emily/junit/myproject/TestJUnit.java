package emily.junit.myproject;

import static org.junit.Assert.*;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;

import junit.framework.JUnit4TestAdapter;
import junit.textui.TestRunner;

public class TestJUnit {
	
	MyApp obj;

	@BeforeClass
	public static void setUpBeforeClass() throws Exception {
	}

	@AfterClass
	public static void tearDownAfterClass() throws Exception {
	}

	@Before
	public void setUp() throws Exception {
	}

	@After
	public void tearDown() throws Exception {
	}

	@Test
	public void test() {
		
		//fail("Not yet implemented");
		assertEquals(5, new MyApp().add(3,2));
	}
	
	@Test(expected = ArithmeticException.class)
	public void testDivide(){
		obj.divide(5,0);
	}
	
	public static void main(){
		TestRunner.run(new JUnit4TestAdapter (TestJUnit.class));
	}

}
